class Admin::UserController < Admin::BaseController
  before_action :set_uesr, only: [:destroy]

  def index
    @users = User.all.order(created_at: :asc)
  end

  def show
    @user = User.find params[:id]
  end

  def new
    @user = User.new
  end

  def edit
    @user = User.find params[:id]
  end

  def create
    # 先判断是否邮箱已经被注册了
    @user = User.new(user_params)
    if @user.save
      flash["success"] = 'create success!'
      redirect_to admin_users_path
    else
      flash["danger"] = '参数错误'
      render :new
    end
  end

  def update
    if params[:user][:password].blank?
      user_params.delete(:password)
      user_params.delete(:password_confirmation)
    end

    @user = User.find params[:id]
    if @user.update user_params
      flash["success"] = 'update success!'
      redirect_to admin_users_path
    else
      render :edit
    end
  end
  
  def destroy
    if @uesr.group.master
      flash['danger'] = 'master role can not delete!'
    else
      @uesr.delete
      flash['success'] = 'delete success'
    end
    redirect_to admin_users_path
  end


  private
  def set_uesr
    @uesr = User.find params[:id]
  end

  def user_params
    new_params = params.require(:user).permit!
    # if params[:user][:authority].present?
    #   new_params[:authority]= params[:user][:authority]
    # end
    # new_params
  end
end
