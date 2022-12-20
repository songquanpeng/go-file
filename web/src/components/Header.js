import React, { useContext, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../context/User';
import { API, showSuccess } from '../helpers';
import '../index.css';

const Header = () => {
  const [userState, userDispatch] = useContext(UserContext);
  let navigate = useNavigate();

  const [showNavbar, setShowNavbar] = useState(false);

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
              href='/'
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
              <a className='navbar-item' href='/'>
                首页
              </a>
              <a className='navbar-item' href='/explorer'>
                文件
              </a>
              <a className='navbar-item' href='/image'>
                图床
              </a>
              <a className='navbar-item' href='/video'>
                视频
              </a>
              <a className='navbar-item' href='/help'>
                帮助
              </a>
            </div>
            <div className='navbar-end'>
              {userState.user ? (
                <div className='navbar-item has-dropdown is-hoverable'>
                  <a className='navbar-link'>{userState.user.username}</a>
                  <div className='navbar-dropdown'>
                    <a className='navbar-item' href='/setting'>
                      设置
                    </a>
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
