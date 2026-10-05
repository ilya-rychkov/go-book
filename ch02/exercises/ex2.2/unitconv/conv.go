package unitconv

func FToM(f Foot) Metre { return Metre(f * 0.3048) }

func MtoF(m Metre) Foot { return Foot(m / 0.3048) }

func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

func KToP(k Kilogram) Pound { return Pound(k * 2.20462262185) }
