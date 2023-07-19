package util

import (
	"bytes"
	"common/proto/comm"
	user "common/proto/user"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	TimestampTolerance = 5 * time.Minute
)

var salt = []byte{109, 99, 97, 114, 100, 50, 48, 50, 50}

func HashSalt(s string) string {
	hash := md5.Sum(append([]byte(s), salt...))
	return fmt.Sprintf("%x", hash)
}

type httpResp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func GetSignature() string {
	t := time.Now().UnixNano()
	t /= int64(time.Millisecond)
	rng := rand.New(rand.NewSource(t / 60))
	num := 10000 + rng.Intn(100000000-10000)
	h := hmac.New(sha256.New, []byte("oconn1020"))
	h.Write([]byte(fmt.Sprintf("%d;%d", t, num)))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%d;%d;%s", t, num, sign)
}

func CheckSignature(sig string) bool {
	parts := strings.Split(sig, ";")
	if len(parts) != 3 {
		return false
	}

	ts, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return false
	}
	ts *= int64(time.Millisecond)
	now := time.Now().UnixNano()
	diff := now - ts
	if diff < 0 {
		diff = -diff
	}
	if diff >= int64(TimestampTolerance) {
		sec := int64(time.Second)
		fmt.Printf("CheckSignature timestamp fail: %d, %d, %d", now/sec, ts/sec, diff/sec)
		return false
	}

	h := hmac.New(sha256.New, []byte("oconn1020"))
	h.Write([]byte(parts[0] + ";" + parts[1]))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	if sign != parts[2] {
		return false
	}

	//redisKey := fmt.Sprintf(defs.KeySignature, sign)
	//if util.RedisGet(ThisServer.RedisClient, redisKey) != "" {
	//	return false
	//}

	//util.RedisSet(ThisServer.RedisClient, redisKey, "1", 2*TimestampTolerance)
	return true
}

func HttpPostJson(host, api string, req any, token string) (p []byte, resp *http.Response, err error) {
	hc := http.Client{}
	url := host + api
	p, err = json.Marshal(req)
	if err != nil {
		return
	}
	httpRequest, err := http.NewRequest("POST", url, bytes.NewReader(p))
	if err != nil {
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("X-Client-Type", "app")
	httpRequest.Header.Set("X-Version", "0.1.1")
	httpRequest.Header.Set("X-Platform", "iOS")
	httpRequest.Header.Set("X-Signature", GetSignature())
	if token != "" {
		httpRequest.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err = hc.Do(httpRequest)
	if err != nil {
		return
	}
	//defer resp.Body.Close()
	p, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return p, resp, nil
}

func HttpPostProto(host, api, token string, pb proto.Message) (p []byte, resp *http.Response, err error) {
	hc := http.Client{}
	url := host + api
	b, err := proto.Marshal(pb)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err = hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
	}
	p, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return p, resp, nil
}

func Otp(host, mail string) (*http.Response, error) {
	api := "/api/user/otp"
	body, resp, err := HttpPostJson(host, api, &user.OtpReq{
		Email:   mail,
		OtpType: 1,
	}, "")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
		return nil, errors.New(resp.Status)
	}
	fmt.Println(string(body))
	//fmt.Println(resp.Header)
	return resp, nil
}

func Register(host, name, mail, passwd, otp string) (*http.Response, error) {
	api := "/api/user/register"
	body, resp, err := HttpPostJson(host, api, &user.UserReq{
		ClientInfo: nil,
		Email:      mail,
		Password:   HashSalt(passwd),
		Otp:        otp, //Otp后从邮箱里查看 dev用123456不需激活
		Name:       name,
	}, "")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
		return nil, errors.New(resp.Status)
	}
	//fmt.Println(string(body))
	var jsonResp httpResp[any]
	if err = json.Unmarshal(body, &jsonResp); err != nil {
		return nil, err
	}
	if jsonResp.Code != 0 {
		return nil, errors.New(jsonResp.Msg)
	}
	//fmt.Println(resp.Header)
	return resp, nil
}

func QuickRegister(host, mail, passwd string) (string, error) {
	api := "/api/user/quick-register"
	body, resp, err := HttpPostJson(host, api, &user.QuickRegisterReq{
		Email:    mail,
		Password: HashSalt(passwd)},
		"")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 0 && resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}
	var jsonResp httpResp[user.UserResp]
	if err = json.Unmarshal(body, &jsonResp); err != nil {
		return "", err
	}
	if jsonResp.Code != 0 {
		return "", errors.New(jsonResp.Msg)
	}
	fmt.Printf("email:%s auth:%s\n", mail, jsonResp.Data.Authorization)
	return jsonResp.Data.Authorization, nil
}

func Login(host, mail, passwd string) (string, error) {
	api := "/api/user/login"
	body, resp, err := HttpPostJson(host, api, &user.UserReq{
		Email:    mail,
		Password: HashSalt(passwd)},
		"")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 0 && resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}
	var jsonResp httpResp[user.UserResp]
	if err = json.Unmarshal(body, &jsonResp); err != nil {
		return "", err
	}
	if jsonResp.Code != 0 {
		return "", errors.New(jsonResp.Msg)
	}
	fmt.Printf("email:%s auth:%s\n", mail, jsonResp.Data.Authorization)
	return jsonResp.Data.Authorization, nil
}

func PlayerList(host, token, Host string) (int64, error) {
	api := "/api/player/get-info"

	body, _, err := HttpPostJson(host, api, &user.PlayerIdReq{
		Host:   Host,
		UserId: 0,
	}, token)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	var jsonResp httpResp[comm.UserLoginResp]
	if err = json.Unmarshal(body, &jsonResp); err != nil {
		return 0, errors.New(err.Error() + "|body:" + string(body))
	}
	if jsonResp.Code != 0 {
		return 0, errors.New(jsonResp.Msg)
	}
	return jsonResp.Data.PlayerId, nil
}

func SetCharacterCreator(host, token string, plrId int64, req *comm.SetCharacterCreatorReq) error {
	req.Source = comm.SetCharacterCreatorReq_CreatePlayer
	req.PlayerId = plrId

	api := "/api/player/character-creator"
	_, _, err := HttpPostJson(host, api, req, token)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDisplayName(host, token string, plrId int64, newName string) error {
	api := "/api/player/change-display-name"
	_, _, err := HttpPostJson(host, api, &comm.PlayerChangeDisplayNameReq{
		PlayerId: plrId,
		Name:     newName,
		Source:   comm.PlayerChangeDisplayNameReq_CreatePlayer,
	}, token)

	if err != nil {
		return err
	}

	return nil
}

func CreatPlayer(host, token, Hostid string) (int64, error) {
	api := "/api/player/create"
	body, _, err := HttpPostJson(host, api, &comm.PlayerCreateReq{
		UserId:      0,
		DisplayName: token[:6],
		PlayerName:  "",
		Host:        Hostid,
		Region:      "Asia/Shanghai",
		AppInfo: &comm.AppInfo{
			DistinctId: "go_test_client" + token,
			DeviceId:   "go_test_client",
			Platform:   "win",
			Version:    0,
		},
	}, token)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	var jsonResp httpResp[comm.PlayerCreateResp]
	if err = json.Unmarshal(body, &jsonResp); err != nil {
		return 0, err
	}
	if jsonResp.Code != 0 {
		return 0, errors.New(jsonResp.Msg)
	}

	UpdateDisplayName(host, token, jsonResp.Data.PlayerInfo.Id, "testone")
	SetCharacterCreator(host, token, jsonResp.Data.PlayerInfo.Id, &comm.SetCharacterCreatorReq{})
	return jsonResp.Data.PlayerInfo.Id, nil
}

func AutoLogin(host, mail, passwd, hostId string) (string, error) {
	author, err := Login(host, mail, passwd)
	if err != nil {
		fmt.Println(err)
		//注册
		author, err = QuickRegister(host, mail, passwd)
		if err != nil {
			return "", err
		}
	}
	playerId, err := PlayerList(host, author, hostId)
	if err != nil {
		return "", err
	}
	if playerId == 0 {
		//创建角色
		playerId, err = CreatPlayer(host, author, hostId)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		fmt.Println(playerId)
	}
	return author, nil
}

func BatchRegister(host, namePre, mailFmt string, from, to int) (list [][]string) {
	for i := from; i <= to; i++ {
		mail := fmt.Sprintf(mailFmt, i)
		passwd := fmt.Sprintf("%d", 100000+rand.Intn(899999))
		var name string
		str := strings.Split(mail, "@")
		if len(str) > 0 {
			name = namePre + str[0]
		}

		//fmt.Println("reg:", name, mail, passwd)
		_, err := Register(host, name, mail, passwd, "123456")
		if err != nil {
			fmt.Println(err)
		} else {
			//fmt.Println(mail, passwd)
			list = append(list, []string{mail, passwd})
		}
	}
	return list
}
