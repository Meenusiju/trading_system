import React, { useMemo } from "react";
import ReactApexChart from "react-apexcharts";
import { ApexOptions } from "apexcharts";
import { Candlestick } from "../types";

interface ChartProps {
  symbol: string;
  candlesticks: Candlestick[];
}

const Chart: React.FC<ChartProps> = ({ symbol, candlesticks }) => {
  const filteredCandlesticks = useMemo(() => {
    return candlesticks.filter((c) => c.Symbol === symbol);
  }, [symbol, candlesticks]);

  const options: ApexOptions = {
    chart: {
      type: "candlestick",
      height: 350,
    },
    plotOptions: {
      bar: { horizontal: false, borderRadius: 10, columnWidth: "45%" },
      candlestick: {
        wick: {
          useFillColor: true,
        },
      },
    },
    title: {
      text: `${symbol} Price Chart`,
      align: "left",
    },
    xaxis: {
      type: "datetime",
      tickPlacement: "between",
    },
    yaxis: {
      tooltip: {
        enabled: true,
      },
    },
  };

  const series = [
    {
      data: filteredCandlesticks.map((c) => ({
        x: new Date(c.Timestamp),
        y: [c.Open, c.High, c.Low, c.Close],
      })),
    },
  ];

  return (
    <ReactApexChart
      options={options}
      series={series}
      type="candlestick"
      height={350}
    />
  );
};

export default Chart;
