package band

import (
	"time"

	"github.com/brocaar/lorawan"
)

type kr920FskBand struct {
	band
}

func (b *kr920FskBand) Name() string {
	return "KR920_FSK"
}

func (b *kr920FskBand) GetDefaults() Defaults {
	return Defaults{
		RX2Frequency:     921900000,
		RX2DataRate:      0,
		MaxFCntGap:       16384,
		ReceiveDelay1:    time.Second,
		ReceiveDelay2:    time.Second * 2,
		JoinAcceptDelay1: time.Second * 5,
		JoinAcceptDelay2: time.Second * 6,
	}
}

func (b *kr920FskBand) GetDownlinkTXPower(freq int) int {
	switch freq {
	case 920900000, 921100000, 921300000, 921500000, 921700000, 921900000, 922100000, 922300000, 922500000, 922700000, 922900000, 923100000, 923300000:
		return 23 // ~200mW
	case 917300000, 917900000, 918500000, 919100000, 919700000, 920300000, 920700000:
		return 10 // ~10mW
	default:
		return 4  // ~3mW
	}
}

func (b *kr920FskBand) GetDefaultMaxUplinkEIRP() float32 {
	return 14 // ~25mW
}

func (b *kr920FskBand) GetPingSlotFrequency(lorawan.DevAddr, time.Duration) (int, error) {
	return 923100000, nil
}

func (b *kr920FskBand) GetRX1ChannelIndexForUplinkChannelIndex(uplinkChannel int) (int, error) {
	return uplinkChannel, nil
}

func (b *kr920FskBand) GetRX1FrequencyForUplinkFrequency(uplinkFrequency int) (int, error) {
	return uplinkFrequency, nil
}

func (b *kr920FskBand) ImplementsTXParamSetup(protocolVersion string) bool {
	return false
}

func newKR920FskBand(repeaterCompatible bool) (Band, error) {
	b := kr920FskBand{
		band: band{
			supportsExtraChannels: true,
			dataRates: map[int]DataRate{
				0: {Modulation: LoRaModulation, SpreadFactor: 12, Bandwidth: 125, uplink: true, downlink: true},
				1: {Modulation: LoRaModulation, SpreadFactor: 11, Bandwidth: 125, uplink: true, downlink: true},
				2: {Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 125, uplink: true, downlink: true},
				3: {Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 125, uplink: true, downlink: true},
				4: {Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 125, uplink: true, downlink: true},
				5: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 125, uplink: true, downlink: true},
				6: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 250, uplink: true, downlink: true},
				7: {Modulation: FSKModulation, BitRate: 50000, uplink: true, downlink: true},
			},
			rx1DataRateTable: map[int][]int{
				0: {0, 0, 0, 0, 0, 0},
				1: {1, 0, 0, 0, 0, 0},
				2: {2, 1, 0, 0, 0, 0},
				3: {3, 2, 1, 0, 0, 0},
				4: {4, 3, 2, 1, 0, 0},
				5: {5, 4, 3, 2, 1, 0},
				6: {6, 5, 4, 3, 2, 1},
				7: {7, 6, 5, 4, 3, 2},
			},
			txPowerOffsets: []int{
				0,
				-2,
				-4,
				-6,
				-8,
				-10,
				-12,
				-14,
			},
			uplinkChannels: []Channel{
				{Frequency: 922100000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 922300000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 922500000, MinDR: 0, MaxDR: 5, enabled: true},
			},

			downlinkChannels: []Channel{
				{Frequency: 922100000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 922300000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 922500000, MinDR: 0, MaxDR: 5, enabled: true},
			},
		},
	}

	if repeaterCompatible {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			latest: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // Added EU868's DR6 and DR7 into LoRaWAN 1.0.2B, 1.1.0A
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
				},
			},
		}
	} else {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			latest: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // Added EU868's DR6 and DR7 LoRaWAN 1.0.2B, 1.1.0A
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					6: {M: 250, N: 242},
					7: {M: 250, N: 242},
				},
			},
		}
	}

	return &b, nil
}
