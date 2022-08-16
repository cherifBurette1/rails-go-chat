class Application < ApplicationRecord
  has_secure_token :access_token

  has_many :chats, dependent: :destroy

  validates :name, presence: true, length: { in: 1..80 }
end
