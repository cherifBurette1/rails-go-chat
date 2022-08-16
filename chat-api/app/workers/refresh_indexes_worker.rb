class RefreshIndexesWorker
  include Sidekiq::Worker
  sidekiq_options queue: :index

  def perform
    unless Message.__elasticsearch__.index_exists?
      Message.__elasticsearch__.create_index!
    end
    Message.import
  end
end