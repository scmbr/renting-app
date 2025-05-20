import React from 'react';

const FilterPanel = ({ selectedCity, onChange }) => {
  const handleSearch = () => {
    if (!selectedCity.trim()) return;
    onChange({ city: selectedCity });
  };

  return (
    <div className="mb-4 flex gap-2">
      <button
        onClick={handleSearch}
        className="bg-blue-500 text-white px-4 py-2 rounded"
      >
        Поиск
      </button>
    </div>
  );
};

export default FilterPanel;
