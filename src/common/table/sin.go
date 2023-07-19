package table

import "math"

var sinTable = []float32{
	0, 0.017452406, 0.034899496, 0.052335955, 0.06975647, 0.087155744, 0.104528464, 0.12186934, 0.1391731, 0.15643446,
	0.17364818, 0.190809, 0.20791169, 0.22495104, 0.24192189, 0.25881904, 0.27563736, 0.2923717, 0.309017, 0.32556814,
	0.34202012, 0.35836795, 0.37460658, 0.39073113, 0.40673664, 0.42261824, 0.43837115, 0.4539905, 0.46947154, 0.4848096,
	0.5, 0.5150381, 0.52991927, 0.54463905, 0.5591929, 0.57357645, 0.58778524, 0.601815, 0.61566144, 0.6293204,
	0.6427876, 0.656059, 0.6691306, 0.6819984, 0.69465834, 0.70710677, 0.7193398, 0.7313537, 0.7431448, 0.75470954,
	0.76604444, 0.7771459, 0.7880107, 0.7986355, 0.809017, 0.81915206, 0.82903755, 0.83867055, 0.8480481, 0.8571673,
	0.8660254, 0.8746197, 0.88294756, 0.8910065, 0.89879405, 0.90630776, 0.9135454, 0.92050487, 0.92718387, 0.9335804,
	0.9396926, 0.94551855, 0.9510565, 0.9563047, 0.9612617, 0.9659258, 0.9702957, 0.97437006, 0.97814757, 0.98162717,
	0.9848077, 0.98768836, 0.99026805, 0.99254614, 0.9945219, 0.9961947, 0.9975641, 0.9986295, 0.99939084, 0.9998477,
	1, 0.9998477, 0.99939084, 0.9986295, 0.9975641, 0.9961947, 0.9945219, 0.99254614, 0.99026805, 0.98768836,
	0.9848077, 0.98162717, 0.9781476, 0.97437006, 0.9702957, 0.9659258, 0.9612617, 0.9563048, 0.95105654, 0.9455186,
	0.9396926, 0.93358046, 0.92718387, 0.92050487, 0.9135455, 0.9063078, 0.89879405, 0.8910065, 0.8829476, 0.8746197,
	0.86602545, 0.8571673, 0.8480481, 0.8386706, 0.8290376, 0.81915206, 0.809017, 0.79863554, 0.7880108, 0.777146,
	0.7660445, 0.7547096, 0.74314487, 0.7313537, 0.71933985, 0.7071068, 0.6946584, 0.6819984, 0.6691306, 0.6560591,
	0.64278764, 0.62932044, 0.6156615, 0.60181504, 0.5877853, 0.57357645, 0.55919296, 0.54463905, 0.5299193, 0.51503813,
	0.50000006, 0.48480967, 0.4694716, 0.45399055, 0.43837118, 0.4226183, 0.40673667, 0.3907312, 0.37460664, 0.35836798,
	0.34202018, 0.3255682, 0.30901703, 0.29237175, 0.2756374, 0.2588191, 0.24192195, 0.2249511, 0.20791174, 0.19080904,
	0.17364822, 0.15643452, 0.13917315, 0.12186939, 0.10452852, 0.0871558, 0.06975652, 0.052336007, 0.03489955, 0.01745246,
	5.3589794e-08, -0.017452352, -0.034899443, -0.052335903, -0.06975642, -0.087155685, -0.104528405, -0.12186929, -0.13917305, -0.15643442,
	-0.17364812, -0.19080894, -0.20791164, -0.224951, -0.24192184, -0.25881898, -0.2756373, -0.29237166, -0.30901694, -0.3255681,
	-0.3420201, -0.3583679, -0.37460655, -0.39073107, -0.40673658, -0.4226182, -0.4383711, -0.45399046, -0.4694715, -0.48480958,
	-0.49999994, -0.515038, -0.5299192, -0.544639, -0.55919284, -0.5735764, -0.5877852, -0.601815, -0.61566144, -0.6293203,
	-0.6427876, -0.65605897, -0.66913056, -0.6819983, -0.69465834, -0.7071067, -0.7193397, -0.73135364, -0.74314475, -0.75470954,
	-0.7660444, -0.7771459, -0.7880107, -0.7986355, -0.80901694, -0.819152, -0.82903755, -0.83867055, -0.84804803, -0.85716724,
	-0.8660254, -0.87461966, -0.88294756, -0.89100647, -0.898794, -0.90630776, -0.9135454, -0.9205048, -0.9271838, -0.9335804,
	-0.9396926, -0.94551855, -0.9510565, -0.9563047, -0.9612617, -0.9659258, -0.9702957, -0.97437006, -0.97814757, -0.98162717,
	-0.9848077, -0.9876883, -0.99026805, -0.99254614, -0.9945219, -0.9961947, -0.997564, -0.9986295, -0.99939084, -0.9998477,
	-1, -0.9998477, -0.99939084, -0.9986295, -0.9975641, -0.9961947, -0.9945219, -0.99254614, -0.99026805, -0.98768836,
	-0.9848078, -0.9816272, -0.9781476, -0.97437006, -0.9702957, -0.9659259, -0.9612617, -0.9563048, -0.95105654, -0.9455186,
	-0.9396927, -0.93358046, -0.92718387, -0.92050487, -0.9135455, -0.9063078, -0.8987941, -0.8910066, -0.8829476, -0.8746197,
	-0.86602545, -0.85716736, -0.84804815, -0.8386706, -0.8290376, -0.8191521, -0.80901706, -0.79863554, -0.78801084, -0.77714604,
	-0.7660445, -0.75470966, -0.74314487, -0.73135376, -0.71933985, -0.7071068, -0.69465846, -0.68199843, -0.6691307, -0.6560591,
	-0.6427877, -0.62932044, -0.61566156, -0.6018151, -0.5877853, -0.5735765, -0.55919296, -0.5446391, -0.5299193, -0.51503813,
	-0.50000006, -0.4848097, -0.46947166, -0.45399058, -0.43837124, -0.42261836, -0.40673673, -0.39073122, -0.3746067, -0.35836804,
	-0.34202024, -0.32556826, -0.3090171, -0.2923718, -0.27563745, -0.25881913, -0.24192199, -0.22495115, -0.20791179, -0.1908091,
	-0.17364828, -0.15643457, -0.13917321, -0.121869445, -0.10452857, -0.08715585, -0.06975658, -0.052336063, -0.034899604, -0.017452514,
}

func Sin(r float32) float32 {
	f := math.Floor(math.Mod(float64(r), 2*360))
	if f < 0 {
		f += 360
	}
	return sinTable[int(f)]
}

func SinByAngle(angle int) float32 {
	angle = angle % 360
	if angle < 0 {
		angle += 360
	}
	return sinTable[angle]
}
