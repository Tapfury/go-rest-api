package conversation

import (
	"net/http"

	messagebird "github.com/messagebird/go-rest-api"
)

type WebhookCreateRequest struct {
	ChannelID string         `json:"channelId"`
	Events    []WebhookEvent `json:"events"`
	URL       string         `json:"url"`
}

// CreateWebhook registers a webhook that is invoked when something interesting
// happens.
func CreateWebhook(c *messagebird.Client, req *WebhookCreateRequest) (*Webhook, error) {
	webhook := &Webhook{}
	if err := request(c, webhook, http.MethodPost, webhooksPath, req); err != nil {
		return nil, err
	}

	return webhook, nil
}

// DeleteWebhook ensures an existing webhook is deleted and no longer
// triggered. If the error is nil, the deletion was successful.
func DeleteWebhook(c *messagebird.Client, id string) error {
	return request(c, nil, http.MethodDelete, webhooksPath+"/"+id, nil)
}

// ListWebhooks gets a collection of webhooks. Pagination can be set in options.
func ListWebhooks(c *messagebird.Client, options *ListOptions) (*WebhookList, error) {
	query := paginationQuery(options)

	webhookList := &WebhookList{}
	if err := request(c, webhookList, http.MethodGet, webhooksPath+"?"+query, nil); err != nil {
		return nil, err
	}

	return webhookList, nil
}

// ReadWebhook gets a single webhook based on its ID.
func ReadWebhook(c *messagebird.Client, id string) (*Webhook, error) {
	webhook := &Webhook{}
	if err := request(c, webhook, http.MethodGet, webhooksPath+"/"+id, nil); err != nil {
		return nil, err
	}

	return webhook, nil
}
