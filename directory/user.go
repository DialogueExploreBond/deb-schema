package directory

// DirectoryUser is the model holding user info on the remote directory.
//
// The model is inspired to the Google Directory Admin user model,
// but some fields are flattened, and it should be treated as a source-agnostic DTO
type DirectoryUser struct {

	// Aliases: Output only. A list of the user's alias email addresses.
	Aliases []string `json:"aliases,omitempty"`

	// Archived: Indicates if user is archived.
	Archived bool `json:"archived,omitempty"`

	// CreationTime: User's G Suite account creation time. (Read-only)
	CreationTime string `json:"creationTime,omitempty"`

	// CustomSchemas: Custom fields of the user.
	OperatorInfo *DirectoryUserOperatorInfo `json:"operatorInfo"`

	DeletionTime string `json:"deletionTime,omitempty"`

	// Etag: Output only. ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: The unique ID for the user. A user `id` can be used as a user
	// request URI's `userKey`.
	Id string `json:"id,omitempty"`

	// IsAdmin: Output only. Indicates a user with super administrator
	// privileges. The `isAdmin` property can only be edited in the Make a
	// user an administrator
	// (/admin-sdk/directory/v1/guides/manage-users.html#make_admin)
	// operation ( makeAdmin
	// (/admin-sdk/directory/v1/reference/users/makeAdmin.html) method). If
	// edited in the user insert
	// (/admin-sdk/directory/v1/reference/users/insert.html) or update
	// (/admin-sdk/directory/v1/reference/users/update.html) methods, the
	// edit is ignored by the API service.
	IsAdmin bool `json:"isAdmin,omitempty"`

	// IsDelegatedAdmin: Output only. Indicates if the user is a delegated
	// administrator. Delegated administrators are supported by the API but
	// cannot create or undelete users, or make users administrators. These
	// requests are ignored by the API service. Roles and privileges for
	// administrators are assigned using the Admin console
	// (https://support.google.com/a/answer/33325).
	IsDelegatedAdmin bool `json:"isDelegatedAdmin,omitempty"`

	// IsEnforcedIn2Sv: Output only. Is 2-step verification enforced
	// (Read-only)
	IsEnforcedIn2Sv bool `json:"isEnforcedIn2Sv,omitempty"`

	// IsEnrolledIn2Sv: Output only. Is enrolled in 2-step verification
	// (Read-only)
	IsEnrolledIn2Sv bool `json:"isEnrolledIn2Sv,omitempty"`

	// Languages: The user's languages. The maximum allowed data size for
	// this field is 1Kb.
	Languages []DirectoryUserLanguage `json:"languages,omitempty"`

	// FamilyName: The user's last name. Required when creating a user
	// account.
	FamilyName string `json:"familyName,omitempty"`

	// FullName: The user's full name formed by concatenating the first and
	// last name values.
	FullName string `json:"fullName,omitempty"`

	// GivenName: The user's first name. Required when creating a user
	// account.
	GivenName string `json:"givenName,omitempty"`

	// OrgUnitPath: The full path of the parent organization associated with
	// the user. If the parent organization is the top-level, it is
	// represented as a forward slash (`/`).
	OrgUnitPath string `json:"orgUnitPath,omitempty"`

	// PrimaryEmail: The user's primary email address. This property is
	// required in a request to create a user account. The `primaryEmail`
	// must be unique and cannot be an alias of another user.
	PrimaryEmail string `json:"primaryEmail,omitempty"`

	// RecoveryEmail: Recovery email of the user.
	RecoveryEmail string `json:"recoveryEmail,omitempty"`

	// RecoveryPhone: Recovery phone of the user. The phone number must be
	// in the E.164 format, starting with the plus sign (+). Example:
	// *+16506661212*.
	RecoveryPhone string `json:"recoveryPhone,omitempty"`

	// Suspended: Indicates if user is suspended.
	Suspended bool `json:"suspended,omitempty"`

	// SuspensionReason: Output only. Has the reason a user account is
	// suspended either by the administrator or by Google at the time of
	// suspension. The property is returned only if the `suspended` property
	// is `true`.
	SuspensionReason string `json:"suspensionReason,omitempty"`

	// ThumbnailPhotoEtag: Output only. ETag of the user's photo (Read-only)
	ThumbnailPhotoEtag string `json:"thumbnailPhotoEtag,omitempty"`

	// ThumbnailPhotoUrl: Output only. Photo Url of the user (Read-only)
	ThumbnailPhotoUrl string `json:"thumbnailPhotoUrl,omitempty"`

	// Groups: the groups the user belongs to
	Groups []DirectoryUserGroup `json:"groups"`

	// Authorities: the permissions granted to the user
	Authorities []DirectoryUserAuthority `json:"authorities"`
}

// DirectoryUserOperatorInfo holds the info regarding BOT Operation.
type DirectoryUserOperatorInfo struct {
	TelegramID             int64 `json:"telegramId,omitempty"`
	ActiveTelegramOperator bool  `json:"activeTelegramOperator,omitempty"`
}

// DirectoryUserLanguage holds info on some user language preference
type DirectoryUserLanguage struct {
	LanguageCode string `json:"languageCode,omitempty"`
	Preference   string `json:"preference,omitempty"`
}

// DirectoryUserAuthority holds a granted authority
type DirectoryUserAuthority struct {
	Code string `json:"code"`
}

// DirectoryUserGroup represents a group the user belongs to
type DirectoryUserGroup struct {
	// AdminCreated: Value is `true` if this group was created by an
	// administrator rather than a user.
	AdminCreated bool `json:"adminCreated,omitempty"`

	// Aliases: List of a group's alias email addresses.
	Aliases []string `json:"aliases,omitempty"`

	// Description: An extended description to help users determine the
	// purpose of a group. For example, you can include information about
	// who should join the group, the types of messages to send to the
	// group, links to FAQs about the group, or related groups. Maximum
	// length is `4,096` characters.
	Description string `json:"description,omitempty"`

	// DirectMembersCount: The number of users that are direct members of
	// the group. If a group is a member (child) of this group (the parent),
	// members of the child group are not counted in the
	// `directMembersCount` property of the parent group.
	DirectMembersCount int64 `json:"directMembersCount,omitempty,string"`

	// Email: The group's email address. If your account has multiple
	// domains, select the appropriate domain for the email address. The
	// `email` must be unique. This property is required when creating a
	// group. Group email addresses are subject to the same character usage
	// rules as usernames, see the help center
	// (https://support.google.com/a/answer/9193374) for details.
	Email string `json:"email,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: The unique ID of a group. A group `id` can be used as a group
	// request URI's `groupKey`.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For Groups resources, the value
	// is `admin#directory#group`.
	Kind string `json:"kind,omitempty"`

	// Name: The group's display name.
	Name string `json:"name,omitempty"`

	// NonEditableAliases: List of the group's non-editable alias email
	// addresses that are outside the account's primary domain or
	// subdomains. These are functioning email addresses used by the group.
	// This is a read-only property returned from the API response for a
	// group. If edited in a group's POST or PUT request, the edit is
	// ignored by the API service.
	NonEditableAliases []string `json:"nonEditableAliases,omitempty"`
}
