import React, { useContext, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { UserContext } from '../context/User';
import { API, showSuccess } from '../helpers';
import '../index.css';
import { SearchContext } from '../context/SearchContext';

const Header = () => {
  const [userState, userDispatch] = useContext(UserContext);
  let navigate = useNavigate();

  const [showNavbar, setShowNavbar] = useState(false);
  const [searchValue, setSearchValue] = useState('');
  const searchContext = useContext(SearchContext);

  const handleSearchChange = (event) => {
    const value = event.target.value;
    setSearchValue(value);
  };

  async function logout() {
    setShowNavbar(false);
    await API.get('/api/user/logout');
    showSuccess('注销成功!');
    userDispatch({ type: 'logout' });
    localStorage.removeItem('user');
    navigate('/login');
  }

  const toggleNavbar = () => {
    setShowNavbar(!showNavbar);
  };

  return (
    <>
      <nav
        className='navbar nav-shadow'
        role='navigation'
        aria-label='main navigation'
        id='nav'
      >
        <div className='container'>
          <div className='navbar-brand'>
            <a
              className='navbar-item is-size-5'
              onClick={() => {
                navigate('/');
              }}
              style={{ fontWeight: 'bold' }}
            >
              Go File
            </a>
            <a
              role='button'
              className={
                'navbar-burger burger' + (showNavbar ? 'is-active' : '')
              }
              aria-label='menu'
              aria-expanded='false'
              data-target='mainNavbar'
              onClick={toggleNavbar}
            >
              <span aria-hidden='true'></span>
              <span aria-hidden='true'></span>
              <span aria-hidden='true'></span>
              <span aria-hidden='true'></span>
              <span aria-hidden='true'></span>
            </a>
          </div>
          <div
            id='mainNavbar'
            className={'navbar-menu' + (showNavbar ? 'is-active' : '')}
          >
            <div className='navbar-start'>
              <Link className='navbar-item' to='/'>
                首页
              </Link>
              <Link className='navbar-item' to='/explorer'>
                文件
              </Link>
              <Link className='navbar-item' to='/image'>
                图床
              </Link>
              <Link to='/video' className='navbar-item'>
                视频
              </Link>
              <Link className='navbar-item' to='/help'>
                帮助
              </Link>
            </div>
            <div className='navbar-end'>
              <div className='navbar-item'>
                <p className='control is-expanded is-light'>
                  <input className='input' type='search' id='searchInput' autoComplete='nope'
                         autoFocus='' value={searchValue} onChange={handleSearchChange} style={{ cursor: 'text' }}
                         placeholder='搜索文件...'
                         onKeyDown={(e) => {
                           if (e.key === 'Enter') {
                             searchContext.setValue(searchValue);
                             console.log("searchContext.setValue ", searchValue);
                             // navigate('/');
                           }
                         }}
                  />
                </p>
              </div>
              {userState.user ? (
                <div className='navbar-item has-dropdown is-hoverable'>
                  <a className='navbar-link'>{userState.user.username}</a>
                  <div className='navbar-dropdown'>
                    <Link className='navbar-item' to='/setting'>
                      设置
                    </Link>
                    <a className='navbar-item' onClick={logout}>
                      注销
                    </a>
                  </div>
                </div>
              ) : (
                <>
                  <div className='navbar-item'>
                    <div className='buttons'>
                      <a className='button is-light'>登录</a>
                    </div>
                  </div>
                  <div className='navbar-item'>
                    <div className='buttons'>
                      <a className='button is-light'>注册</a>
                    </div>
                  </div>
                </>
              )}
            </div>
          </div>
        </div>
      </nav>
    </>
  );
};

export default Header;
