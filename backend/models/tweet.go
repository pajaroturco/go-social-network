package models

/*Tweet estrcutura para un tweet*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
