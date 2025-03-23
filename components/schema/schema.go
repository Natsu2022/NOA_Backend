package schema

type GyroStruct struct {
	Acceleration                   float32 `json:"Acceleration"`
	VelocityAngular                float32 `json:"VelocityAngular"`
	VibrationSpeed                 float32 `json:"VibrationSpeed"`
	VibrationAngle                 float32 `json:"VibrationAngle"`
	VibrationDisplacement          float32 `json:"VibrationDisplacement"`
	VibrationDisplacementHighSpeed float32 `json:"VibrationDisplacementHighSpeed"`
	Frequency                      float32 `json:"Frequency"`
}

type GyroData struct {
	DeviceAddress   string     `json:"DeviceAddress"`
	DateTime        string     `json:"DateTime"`
	TimeStamp       int64      `json:"TimeStamp"`
	X               GyroStruct `json:"X"`
	Y               GyroStruct `json:"Y"`
	Z               GyroStruct `json:"Z"`
	Temperature     float32    `json:"Temperature"`
	ModbusHighSpeed bool       `json:"ModbusHighSpeed"`
}

type PasswordRequest struct {
	Password string `json:"Password"`
	CFP      string `json:"CFP"`
}

type User struct {
	ID       string `bson:"id,omitempty"` // User ID
	Email    string `bson:"email"`        // User email
	Password string `bson:"password"`     // User password
}

type Account struct {
	ID  string `bson:"id,omitempty"` // User ID
	OTP string `bson:"otp"`          // OTP
}
