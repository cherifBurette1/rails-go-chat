class MessageWorker
  include Sidekiq::Worker
  sidekiq_options queue: :message

  def perform(access_token, chat_number, number, body)
    application = Application.find_by!(access_token: access_token)
    chat = application.chats.find_by!(number: chat_number)
    chat.messages.create!(number: number, body: body)
  end
end