#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/runtime:5.0-buster-slim AS base
RUN apt update \
    && apt install -y --no-install-recommends libc-dev \
    && apt-get -y clean \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:5.0-buster-slim AS build
WORKDIR /src
COPY ["app/ConsoleAppEcs/ConsoleAppEcs.csproj", "app/ConsoleAppEcs/"]
COPY ["app/Benchmark.ClientLib/Benchmark.ClientLib.csproj", "app/Benchmark.ClientLib/"]
RUN dotnet restore "app/ConsoleAppEcs/ConsoleAppEcs.csproj"
COPY . .
WORKDIR "/src/app/ConsoleAppEcs"
RUN dotnet build "ConsoleAppEcs.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "ConsoleAppEcs.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "ConsoleAppEcs.dll"]