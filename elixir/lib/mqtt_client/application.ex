defmodule MqttClient.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  def start(_type, _args) do
    children = [
      # Starts a worker, if connection failes will be retarted by supervisor
      {Tortoise.Connection,
       [
         client_id: "elixir_test_client",
         handler: {MqttClientHandler, []},
         server: {Tortoise.Transport.Tcp, host: 'opendata.dekart.com', port: 1945},
         subscriptions: [{"telemetry/transport/+", 0}]
       ]}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: MqttClient.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
