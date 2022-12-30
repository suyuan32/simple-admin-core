package config

type GLobalEnv struct {
	Endtime int64
}

func (g GLobalEnv) NewGlobalEnv() GLobalEnv {
	return GLobalEnv{Endtime: g.Endtime}
}
