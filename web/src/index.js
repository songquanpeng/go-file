import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import 'bulma/css/bulma.min.css';
import './index.css';
import App from './App';
import Header from './components/Header';
import Footer from './components/Footer';
import { UserProvider } from './context/User';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { SearchProvider } from './context/SearchContext';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <UserProvider>
      <SearchProvider>
        <BrowserRouter>
          <Header />
          <App />
          <ToastContainer />
          <Footer />
        </BrowserRouter>
      </SearchProvider>
    </UserProvider>
  </React.StrictMode>
);
