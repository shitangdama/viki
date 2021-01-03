Rails.application.routes.draw do

  
  
  resources :segments
  resources :shares
  resources :spread_trends
  devise_for :users
  post 'authenticate', to: 'authentication#authenticate'
  post 'authed_user', to: 'authentication#authed_user'

  root to: 'admin/message#index' 
  resources :message
  namespace :admin do
    root to: 'message#index' 
    resources :users do
    #   member do
    #     # get :articles_index
    #     # get :user_info
    # end
    end

    resources :segments do

    end

    resources :message
  end


end
