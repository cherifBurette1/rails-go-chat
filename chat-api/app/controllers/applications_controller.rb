class ApplicationsController < ApplicationController

  include Render

  def index
    @applications = Application.all
    render_json @applications
  end

  def create
    @application = Application.new(application_params)
    @application.chat_count = 0
    @application.save
    render_json @application
  end

  def update
    @application = Application.find_by!(access_token: params[:access_token])
    @application.update(application_params)
    render_json @application
  end
  def show
    @application = Application.find_by!(access_token: params[:access_token])
    render_json @application
  end

  private

  def application_params
    params.permit(:name)
  end
end
