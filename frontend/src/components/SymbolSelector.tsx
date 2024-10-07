import React from "react";

interface SymbolSelectorProps {
  symbol: string;
  onSymbolChange: (symbol: string) => void;
}

const SymbolSelector: React.FC<SymbolSelectorProps> = ({
  symbol,
  onSymbolChange,
}) => {
  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedValue = event.target.value;
    onSymbolChange(selectedValue);
  };

  return (
    <select
      className="select-dropdown"
      value={symbol}
      onChange={handleSelectChange}
    >
      <option value="BTCUSDT">BTCUSDT</option>
      <option value="ETHUSDT">ETHUSDT</option>
      <option value="PEPEUSDT">PEPEUSDT</option>
    </select>
  );
};

export default SymbolSelector;
