package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GC4YuMVzIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwNjw4ykxDdhjvmXDUT2vQ93aUu+GmL77Zhg32wLRb64T7Mqo3LOl5mEarQXOjSYtbmUz5ihpuXwxsVRRs1w4V/JR1PqG+eYtF7xXPGCxcmkPbPy/efPz9//rH7ZbHq8UEble4FAg4Mte+lt2SunFrd8k3xlaanWXdGj9UZTuGdzhk1r0ZptKZKvvUR3TmKZciFI6bjxyjUZ8wTP9NsYHH0XeZOn0fK+olxYS2B5tFGXzo2/02LDHy5YmGZlc71f/8vLxE8lWyJsJVsVxZiVI/dZ2O702SfzOj1/8uLrK5IYjzO1L3C4cnnDUw25J81fWB4+u5Dm/felTWWJgWT8i35+bbPZuostWJrZV3VImShu1lvo+8x8EcsbN5eG/qo+IUEThaXbWRXXvSnkLA92uP+m5vDXhf4ahupT8k7U3MjMX/FCTF1aZcH2l1ZvbvVqztthlOy+WkrTochN9UTUuhW3X5rIHXl1bnL55IuOVT+d5s81W7lzjhvrx2+yJ1/8+pJa2Xtqonx+4epDtwX/6++Zezv36VvNop3V1u/f7q55+nb/27dv9j9K/ic88ZP63iqtaGYH3ooyw5vPzCqy/FSMWSw4XLNmNVlMK+rn3L8v1bV2JcvCxrVeMamZohNTva5xXl8Vrr7/+v8WK9er775O+8u8KFz66SzX7zPyr89Jfc52ZPvDkyk2fvFvr/u0S4o7ic6Qv8vt+TI51N7nmabPuraTp41r/zjK7/sv2r1SiXsiR+oGA/baLza1n3++vurxvr9+tvNd1ojT/4/dZ2Zg+P8/wJud456DzA5DJgaG/TwMDLC0x8BwFS3tsSPSHji5Nc1ISgTpRlYT4M3IJMKMSLvIJoPSLgxsawSReFMywijsToEAAYb/jjVMDFgcxsoGkmdiYGLoZGBg2MUE4gECAAD//76sQ/FZAwAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
