package table

import "math"

var cosTable = []float32{
	1, 0.9998477, 0.99939084, 0.9986295, 0.9975641, 0.9961947, 0.9945219, 0.99254614, 0.99026805, 0.98768836,
	0.9848077, 0.98162717, 0.9781476, 0.97437006, 0.9702957, 0.9659258, 0.9612617, 0.9563047, 0.95105654, 0.94551855,
	0.9396926, 0.93358046, 0.92718387, 0.92050487, 0.9135455, 0.9063078, 0.89879405, 0.8910065, 0.8829476, 0.8746197,
	0.8660254, 0.8571673, 0.8480481, 0.83867055, 0.8290376, 0.81915206, 0.809017, 0.79863554, 0.7880108, 0.777146,
	0.76604444, 0.7547096, 0.7431448, 0.7313537, 0.7193398, 0.70710677, 0.6946584, 0.6819984, 0.6691306, 0.656059,
	0.64278764, 0.6293204, 0.6156615, 0.60181504, 0.58778524, 0.57357645, 0.5591929, 0.54463905, 0.52991927, 0.5150381,
	0.5, 0.48480964, 0.46947157, 0.45399052, 0.43837115, 0.42261827, 0.40673667, 0.39073116, 0.3746066, 0.35836798,
	0.34202015, 0.32556817, 0.309017, 0.29237172, 0.2756374, 0.25881907, 0.24192192, 0.22495107, 0.20791171, 0.19080901,
	0.17364821, 0.15643449, 0.13917312, 0.12186937, 0.10452849, 0.08715577, 0.0697565, 0.05233598, 0.03489952, 0.017452434,
	2.6794897e-08, -0.01745238, -0.03489947, -0.05233593, -0.06975645, -0.087155715, -0.104528435, -0.12186932, -0.13917308, -0.15643443,
	-0.17364815, -0.19080897, -0.20791166, -0.22495103, -0.24192187, -0.258819, -0.27563733, -0.29237166, -0.30901697, -0.3255681,
	-0.34202012, -0.35836792, -0.37460655, -0.3907311, -0.4067366, -0.42261824, -0.43837112, -0.45399046, -0.46947154, -0.48480958,
	-0.49999997, -0.5150381, -0.5299192, -0.544639, -0.5591929, -0.5735764, -0.58778524, -0.601815, -0.61566144, -0.6293204,
	-0.6427876, -0.656059, -0.66913056, -0.6819983, -0.69465834, -0.70710677, -0.7193398, -0.7313537, -0.7431448, -0.75470954,
	-0.76604444, -0.7771459, -0.7880107, -0.7986355, -0.80901694, -0.819152, -0.82903755, -0.83867055, -0.8480481, -0.8571673,
	-0.8660254, -0.87461966, -0.88294756, -0.8910065, -0.89879405, -0.90630776, -0.9135454, -0.9205048, -0.92718387, -0.9335804,
	-0.9396926, -0.94551855, -0.9510565, -0.9563047, -0.9612617, -0.9659258, -0.9702957, -0.97437006, -0.97814757, -0.98162717,
	-0.9848077, -0.9876883, -0.99026805, -0.99254614, -0.9945219, -0.9961947, -0.997564, -0.9986295, -0.99939084, -0.9998477,
	-1, -0.9998477, -0.99939084, -0.9986295, -0.9975641, -0.9961947, -0.9945219, -0.99254614, -0.99026805, -0.98768836,
	-0.9848078, -0.98162717, -0.9781476, -0.97437006, -0.9702957, -0.9659258, -0.9612617, -0.9563048, -0.95105654, -0.9455186,
	-0.9396926, -0.93358046, -0.92718387, -0.92050487, -0.9135455, -0.9063078, -0.89879405, -0.8910065, -0.8829476, -0.8746197,
	-0.86602545, -0.8571673, -0.84804815, -0.8386706, -0.8290376, -0.81915206, -0.809017, -0.79863554, -0.7880108, -0.777146,
	-0.7660445, -0.7547096, -0.74314487, -0.73135376, -0.71933985, -0.7071068, -0.6946584, -0.68199843, -0.6691307, -0.6560591,
	-0.64278764, -0.62932044, -0.6156615, -0.6018151, -0.5877853, -0.5735765, -0.55919296, -0.5446391, -0.5299193, -0.51503813,
	-0.50000006, -0.4848097, -0.46947163, -0.45399058, -0.4383712, -0.42261833, -0.4067367, -0.3907312, -0.37460667, -0.358368,
	-0.3420202, -0.32556823, -0.30901706, -0.29237178, -0.27563742, -0.25881913, -0.24192198, -0.22495113, -0.20791176, -0.19080907,
	-0.17364825, -0.15643454, -0.13917318, -0.12186942, -0.10452854, -0.08715582, -0.06975655, -0.052336037, -0.034899577, -0.017452486,
	-8.038469e-08, 0.017452326, 0.034899417, 0.052335873, 0.06975639, 0.08715566, 0.10452838, 0.12186926, 0.13917302, 0.15643439,
	0.17364809, 0.1908089, 0.20791161, 0.22495097, 0.24192181, 0.25881895, 0.27563727, 0.29237163, 0.3090169, 0.32556808,
	0.34202006, 0.35836786, 0.37460652, 0.39073104, 0.40673655, 0.42261818, 0.43837106, 0.45399043, 0.46947148, 0.48480955,
	0.4999999, 0.515038, 0.5299192, 0.54463893, 0.55919284, 0.5735764, 0.5877852, 0.6018149, 0.6156614, 0.6293203,
	0.6427875, 0.65605897, 0.66913056, 0.6819983, 0.6946583, 0.7071067, 0.7193397, 0.73135364, 0.74314475, 0.75470954,
	0.7660444, 0.7771459, 0.7880107, 0.7986354, 0.80901694, 0.819152, 0.82903755, 0.8386705, 0.84804803, 0.85716724,
	0.8660253, 0.87461966, 0.88294756, 0.89100647, 0.898794, 0.90630776, 0.9135454, 0.9205048, 0.9271838, 0.9335804,
	0.93969256, 0.94551855, 0.9510565, 0.9563047, 0.9612617, 0.9659258, 0.9702957, 0.97437006, 0.97814757, 0.98162717,
	0.9848077, 0.9876883, 0.99026805, 0.99254614, 0.99452186, 0.99619466, 0.997564, 0.9986295, 0.99939084, 0.9998477,
}

func Cos(r float32) float32 {
	f := math.Floor(math.Mod(float64(r), 2*360))
	if f < 0 {
		f += 360
	}
	return cosTable[int(f)]
}

func CosByAngle(angle int) float32 {
	angle = angle % 360
	if angle < 0 {
		angle += 360
	}
	return cosTable[angle]
}
