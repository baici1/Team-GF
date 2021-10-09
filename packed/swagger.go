package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBIyfYIZkACIgycDMXlienpqUX6UFovqzg/LzSElYFR1z0hcU3owbzLDgJ7v2uvaU0Wdbd8dmxqsrZfx+Fr83ZEX919e+7GVRw/A/g4ArwncrC7hAoyLlJnehTAyfdX8lHLfodsdc1zt9SUPipOSr+3r/J32rHKPU8r5x3/MMNdUOLa14jlV9Nuu4acmJcUu/7Ak4kOno8YbybMcr8sGVJ867jRNX+Bnn6lSZ16qV3MXCc7xSQXH1JhOP3twslnHOu+LuSw4FT9yxwyyXKD+209/Y8rGouNdpyaF5sunv5Y/PDVHwadyXm/rm2++fB9+OPly2+e9n/8mX/No4ZnbccdEl6l7ig7Iy/GxOaWUZ0ZvUnDTDY2/f68EJFomck2bN3snk2/fD0mfn63coNg6gMNQTana19PqFxoOJIgnxR387HQu3WN89bWHOdfqL5kg7mLfZiizLYyzZ3Lj67reMK8l9k9L/mZf9TMwsWhvSYNUZ4sa66esnpb9th1zjub/zay8/zq9p3T0591J+b84p0frtWGFKV/27TP/emJt9GVj/8LBd37/eqcpTzDq9c1BlfvMbJsSZn0lyPDQCVr1iETd5swtTX5N+fIn+SQ3LluyVTpqA1lZt7dUzpKp1z7+nibyGq5Vt7ZcisuPI1JmR724PuDtyUms/bYr5z5v1VMkfNN/XmvsLY5r66nRh76kmZQVz5z789Zc2t/Tu+aqazd2WEYFelc9i/r08fj62cnlhn9yuctZd4Q+mdZPRMDw///Ad7sHM1LWtYyMjEwHGNnYIClJQaMtMSOSEvw5APSjawmwJuRSYQZkRaRTQalRRjY1ggi8aZMhFHYnQIBAgz/HX2YGLA4jJUNJM/EwMTQycDA0MUE4gECAAD//3HpzK8pAwAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
