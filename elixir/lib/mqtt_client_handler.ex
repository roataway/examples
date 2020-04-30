defmodule MqttClientHandler do
  @moduledoc """
  MqttClientHandler
  """
  use Tortoise.Handler

  def init(args) do
    {:ok, args}
  end

  def connection(_status, state) do
    {:ok, state}
  end

  @doc """
  Get messages from topic /telemetry/transport/+
  """
  def handle_message(["telemetry", "transport", _track], payload, state) do
    # decods json to a map
    _decoded = Jason.decode!(payload)
    {:ok, state}
  end

  @doc """
  Callback to subscribe to a topic
  """
  def subscription(_status, _topic_filter, state) do
    {:ok, state}
  end

  @doc """
  Callback when client is terminated
  """
  def terminate(_reason, _state) do
    :ok
  end
end
