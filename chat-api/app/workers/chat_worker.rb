class ChatWorker
  include Sidekiq::Worker
  sidekiq_options queue: :chat

  def perform(access_token, number)
    application = Application.find_by!(access_token: access_token)
    application.chats.create!(number: number, message_count: 0)
  end
end