class AddChatCountToApplications < ActiveRecord::Migration[6.0]
  def change
    add_column :applications, :chat_count, :integer
  end
end
