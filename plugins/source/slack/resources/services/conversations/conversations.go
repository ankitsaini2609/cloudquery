// Code generated by codegen; DO NOT EDIT.

package conversations

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Conversations() *schema.Table {
	return &schema.Table{
		Name:        "slack_conversations",
		Description: `https://api.slack.com/methods/conversations.list`,
		Resolver:    fetchConversations,
		Multiplex:   client.TeamMultiplex,
		Columns: []schema.Column{
			{
				Name:     "team_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveTeamID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("Created"),
			},
			{
				Name:     "is_open",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsOpen"),
			},
			{
				Name:     "last_read",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastRead"),
			},
			{
				Name:     "unread_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("UnreadCount"),
			},
			{
				Name:     "unread_count_display",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("UnreadCountDisplay"),
			},
			{
				Name:     "is_group",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsGroup"),
			},
			{
				Name:     "is_shared",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsShared"),
			},
			{
				Name:     "is_im",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsIM"),
			},
			{
				Name:     "is_ext_shared",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsExtShared"),
			},
			{
				Name:     "is_org_shared",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsOrgShared"),
			},
			{
				Name:     "is_pending_ext_shared",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsPendingExtShared"),
			},
			{
				Name:     "is_private",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsPrivate"),
			},
			{
				Name:     "is_mpim",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsMpIM"),
			},
			{
				Name:     "unlinked",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Unlinked"),
			},
			{
				Name:     "name_normalized",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NameNormalized"),
			},
			{
				Name:     "num_members",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumMembers"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "user",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "creator",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Creator"),
			},
			{
				Name:     "is_archived",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsArchived"),
			},
			{
				Name:     "topic",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Topic"),
			},
			{
				Name:     "purpose",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Purpose"),
			},
			{
				Name:     "is_channel",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsChannel"),
			},
			{
				Name:     "is_general",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsGeneral"),
			},
			{
				Name:     "is_member",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsMember"),
			},
			{
				Name:     "locale",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Locale"),
			},
		},

		Relations: []*schema.Table{
			ConversationBookmarks(),
			ConversationHistories(),
			ConversationMembers(),
		},
	}
}
