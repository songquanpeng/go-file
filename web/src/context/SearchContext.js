import React, { useState } from 'react';

export const SearchContext = React.createContext('');

export const SearchProvider = ({ children }) => {
  const [searchValue, setSearchValue] = useState('');

  return (
    <SearchContext.Provider value={{ value: searchValue, setValue: setSearchValue }}>
      {children}
    </SearchContext.Provider>
  );
};