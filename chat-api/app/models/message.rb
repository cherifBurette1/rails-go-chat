class Message < ApplicationRecord
  belongs_to :chat, counter_cache: :message_count, touch: true

  validates :number,
            presence: true,
            numericality: { only_integer: true, greater_than_or_equal_to: 1 },
            uniqueness: { scope: :chat_id }

  include Searchable
end
