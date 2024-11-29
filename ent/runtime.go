// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/Jamescog/birth-buddy/ent/schema"
	"github.com/Jamescog/birth-buddy/ent/users"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	usersFields := schema.Users{}.Fields()
	_ = usersFields
	// usersDescTelegramID is the schema descriptor for telegram_id field.
	usersDescTelegramID := usersFields[0].Descriptor()
	// users.TelegramIDValidator is a validator for the "telegram_id" field. It is called by the builders before save.
	users.TelegramIDValidator = usersDescTelegramID.Validators[0].(func(int) error)
	// usersDescIsPremium is the schema descriptor for is_premium field.
	usersDescIsPremium := usersFields[3].Descriptor()
	// users.DefaultIsPremium holds the default value on creation for the is_premium field.
	users.DefaultIsPremium = usersDescIsPremium.Default.(bool)
	// usersDescBirthdayCount is the schema descriptor for birthday_count field.
	usersDescBirthdayCount := usersFields[4].Descriptor()
	// users.DefaultBirthdayCount holds the default value on creation for the birthday_count field.
	users.DefaultBirthdayCount = usersDescBirthdayCount.Default.(int)
}
