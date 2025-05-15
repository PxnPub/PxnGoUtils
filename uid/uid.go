package uid;

import(
	Fmt    "fmt"
	IO     "io"
	BufIO  "bufio"
	Sync   "sync"
	Time   "time"
	Errors "errors"
	Rand   "crypto/rand"
);



const MaxTimestamp int64 = 0x8ffffff_ffffffff;
const MaxCounter   int64 = 0b0011_1111;
const MaxID        int64 = 0b11;



// 6 bytes unix timestamp
// 1 byte  entropy from /dev/urandom
// 1 byte  counter & id
type UID64 uint64;



type Generator struct {
	ID        uint8
	RNG       IO.Reader
	Count     uint8
	Timestamp int64
	Mut       Sync.Mutex
}



func New(id int) *Generator {
	return &Generator{
		ID:  uint8(id % 4),
		RNG: BufIO.NewReaderSize(Rand.Reader, 16),
	};
}



func (gen *Generator) Next() (UID64, error) {
	gen.Mut.Lock();
	defer gen.Mut.Unlock();
	return gen.NextUnsafe();
}

func (gen *Generator) NextUnsafe() (UID64, error) {
	now := Time.Now().UnixMilli();
	if now == gen.Timestamp {
		if gen.Count == 0xff {
			return 0, errors.New("Gen max ratio exceeded");
		}
		gen.Count++;
	} else {
		gen.Count = 0;
	}
	b := [1]byte{};
	if _, err := gen.RNG.Read(b[:]); err != nil {
		return 0, err;
	}
	rnd := b[0];
	return ToUID(now, rnd, gen.ID, gen.Count);
}



func NewUID(now int64, rnd byte, id int, count uint8) (UID, error) {
	if now > TimestampMax { return 0, Fmt.Error("Timestamp can't exceed 0x8ffffff_ffffffff"); }
	if id > MaxID         { return 0, Fmt.Error("ID can't exceed %d", MaxID); }
	if count > MaxCounter { return 0, Fmt.Error("Counter can't exceed %d", MaxCounter); }
	return UID(
		(uint64(timestamp) << 16) +
		(uint64(entropy)   <<  8) +
		uint64(generatorID <<  6) +
		uint64(counter)
	);
}



func Parse(str string) (UID, error) {
	uid, err := FromBase36(str);
	return UID(uid), err;
}

func (uid UID) ToString() string {
	return ToBase36(uid);
}



func FromInt(val int64) (UID, error) {
	return UID(val), nil;
}

func (uid UID) ToInt64() int64 {
	return int64(uid);
}



func (uid UID) GetTimestamp() int64 {
	return int64((uid & 0xffffffff_ffff0000) >> 16);
}

func (uid UID) GetEntropy() uint8 {
	return uint8((uid & 0xff00) >> 8);
}

func (uid UID) GetCounter() uint8 {
	return uint8(uid & MaxCounter);
}

func (uid UID) GetID() uint8 {
	return uint8((uid & 0b1100_0000) >> 6);
}
