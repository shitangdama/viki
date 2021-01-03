class Admin::SegmentsController < Admin::BaseController
  before_action :set_segment, only: [:show, :edit, :update, :destroy]

  def index
    @segments = Segment.all
  end

  def swt_index
    @segments = Segment.all
  end

  def show
  end

  def new
    @segment = Segment.new
  end

  def edit
  end

  def create
    @segment = Segment.new(segment_params)
  end

  def update
  end

  def destroy
    @segment.destroy
  end

  private
    def set_segment
      @segment = Segment.find(params[:id])
    end

    def segment_params
      params.fetch(:segment, {})
    end
end
