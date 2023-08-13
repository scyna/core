package base

import (
	"encoding/json"
	"strconv"
	"sync"

	scyna_const "github.com/scyna/core/const"
	scyna_proto "github.com/scyna/core/proto/generated"
)

type Settings struct {
	data   map[string]string /*cache*/
	mutex  sync.Mutex
	conn   *Nats
	module string
}

func (s *Settings) Init(module string, conn *Nats) {
	s.data = make(map[string]string)
}

func (s *Settings) Remove(key string) bool {
	request := scyna_proto.RemoveSettingRequest{Module: s.module, Key: key}
	if err := s.conn.SendRequest(scyna_const.SETTING_REMOVE_URL, &request); err.Code() == 0 {
		s.removed(key)
		return true
	}
	return false
}

func (s *Settings) Write(key string, value string) bool {
	request := scyna_proto.WriteSettingRequest{Module: s.module, Key: key, Value: value}
	if err := s.conn.SendRequest(scyna_const.SETTING_WRITE_URL, &request); err.Code() == 0 {
		s.updated(key, value)
		return true
	}
	return false
}

func (s *Settings) ReadString(key string) (bool, string) {
	/*from cache*/
	s.mutex.Lock()
	if val, ok := s.data[key]; ok {
		s.mutex.Unlock()
		return true, val
	}
	s.mutex.Unlock()

	/*from manager*/
	request := scyna_proto.ReadSettingRequest{Module: s.module, Key: key}
	var response scyna_proto.ReadSettingResponse
	if err := s.conn.SendAndReceive(scyna_const.SETTING_READ_URL, &request, &response); err.Code() == 0 {
		s.updated(key, response.Value)
		return true, response.Value
	}
	return false, ""
}

func (s *Settings) ReadInt(key string) (bool, int) {
	if ok, val := s.ReadString(key); ok {
		if i, err := strconv.Atoi(val); err != nil {
			return false, 0
		} else {
			return true, i
		}
	}
	return false, 0
}

func (s *Settings) ReadBool(key string) (bool, bool) {
	/*TODO*/
	return false, false
}

func (s *Settings) ReadObject(key string, value interface{}) bool {
	if ok, val := s.ReadString(key); ok {
		if err := json.Unmarshal([]byte(val), value); err != nil {
			//Session.Error("ReadObjectSetting: " + err.Error())
			return false
		}
		return true
	}
	return false
}

func (s *Settings) updated(key string, value string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = value
}

func (s *Settings) removed(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, key)
}
