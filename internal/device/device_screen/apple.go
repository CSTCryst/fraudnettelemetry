package device_screen

import "github.com/colduction/randomizer"

var (
	appleSmartPhoneLogicalRes  [20][2]uint32 = [20][2]uint32{{320, 480}, {640, 960}, {640, 1136}, {750, 1334}, {1080, 2340}, {1125, 2436}, {1170, 2532}, {1080, 1920}, {828, 1792}, {1242, 2688}, {1284, 2778}, {1242, 2208}, {1488, 2266}, {768, 1024}, {1536, 2048}, {1620, 2160}, {1640, 2360}, {1668, 2224}, {1668, 2388}, {2048, 2732}}
	appleSmartPhonePhysicalRes [20][2]uint32 = [20][2]uint32{{320, 480}, {320, 480}, {320, 568}, {375, 667}, {375, 812}, {375, 812}, {390, 844}, {414, 736}, {414, 896}, {414, 896}, {428, 926}, {476, 847}, {744, 1133}, {768, 1024}, {768, 1024}, {810, 1080}, {820, 1180}, {834, 1112}, {834, 1194}, {1024, 1366}}
	appleSmartPhoneColorDepth  [4]uint8      = [4]uint8{24, 30, 36, 48}
)

func (m *screen) SetAppleSmartphone() {
	i := randomizer.IntInterval(0, len(appleSmartPhoneLogicalRes))
	m.LogicalResolution = appleSmartPhoneLogicalRes[i]
	m.PhysicalResolution = appleSmartPhonePhysicalRes[i]
	m.ColorDepth = appleSmartPhoneColorDepth[randomizer.IntInterval(0, len(appleSmartPhoneColorDepth))]
}
