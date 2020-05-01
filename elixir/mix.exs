defmodule MqttClient.MixProject do
  use Mix.Project

  def project do
    [
      app: :mqtt_client,
      version: "0.1.0",
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger],
      mod: {MqttClient.Application, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:tortoise, "~> 0.9.4"},
      {:jason, "~> 1.2"}
      # {:dep_from_hexpm, "~> 0.3.0"},
      # {:dep_from_git, git: "https://github.com/elixir-lang/my_dep.git", tag: "0.1.0"}
    ]
  end
end
