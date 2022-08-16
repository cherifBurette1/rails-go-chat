class ChatsController < ApplicationController

  include Render
  before_action :set_application

  def index
    @chats = @application.chats.all
    render_json @chats
  end

  def show
    @chat = @application.chats.find_by!(number: params[:number]) # TODO: Handle application not found
  end

  private

  def chat_params
    params.permit(:number)
  end

  def set_application
    @application = Application.find_by!(access_token: params[:application_access_token])
  end
end
