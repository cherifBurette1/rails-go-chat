class ApplicationController < ActionController::API
  include ActionController::MimeResponds

  rescue_from StandardError do |e|
    if e.is_a? ActiveRecord::RecordNotFound
      render :json => {:error => "404: Record not found"}.to_json, :status => 404
    else
      render :json => {:error => "500: Internal error"}.to_json, :status => 500  end
    end
end
