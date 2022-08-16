Sidekiq.configure_server do |config|
 config.redis = { host:'redis', port: 6379, db: 0 }
end

Sidekiq.configure_client do |config|
 config.redis = { host:'redis', port: 6379, db: 0 }
end
