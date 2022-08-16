class AddMessageCountToChats < ActiveRecord::Migration[6.0]
  def change
    add_column :chats, :message_count, :integer
  end
end
