package data

import (
	log "github.com/sirupsen/logrus"
)

// GetString returns the string value for key.
func (ds *Datastore) GetString(key string) (string, error) {
	configEntry, err := ds.Get(key)
	if err != nil {
		return "", err
	}
	return configEntry.getString()
}

// SetString will set the string value for the desired key.
func (ds *Datastore) SetString(key string, value string) error {
	configEntry := ConfigEntry{key, value}
	return ds.Save(configEntry)
}

// GetNumber returns the numeric value for key.
func (ds *Datastore) GetNumber(key string) (float32, error) {
	configEntry, err := ds.Get(key)
	if err != nil {
		return 0, err
	}
	return configEntry.getNumber()
}

// SetNumber will set the numeric value for the desired key.
func (ds *Datastore) SetNumber(key string, value float32) error {
	configEntry := ConfigEntry{key, value}
	return ds.Save(configEntry)
}

// GetBool returns the boolean value for key.
func (ds *Datastore) GetBool(key string) (bool, error) {
	configEntry, err := ds.Get(key)
	if err != nil {
		return false, err
	}
	return configEntry.getBool()
}

// SetBool will set the boolean value for the desired key.
func (ds *Datastore) SetBool(key string, value bool) error {
	configEntry := ConfigEntry{key, value}
	return ds.Save(configEntry)
}

// Set will set the value for the desired key.
func (ds *Datastore) Set(key string, value interface{}) error {
	switch value.(type) {
	case int, float32:
		return ds.SetNumber(key, value.(float32))

	case string:
		return ds.SetString(key, value.(string))

	case bool:
		return ds.SetBool(key, value.(bool))
	}

	log.Panicf("Unable to store value of type %T", value)

	return nil
}
