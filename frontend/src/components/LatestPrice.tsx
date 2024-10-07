import React from "react";

interface LatestPriceProps {
  price: number | null;
  color: string;
  chartTime: string;
}

const LatestPrice: React.FC<LatestPriceProps> = ({
  price,
  color,
  chartTime,
}) => {
  return (
    <div className="latest-price">
      <div style={{ color: color }}>
        <span>Latest Price: {price !== null ? price : "N/A"}</span>
      </div>
      <div>
        <p>Last Updated: {chartTime}</p>
        <p>Chart Frame: 1m</p>
      </div>
    </div>
  );
};

export default LatestPrice;
