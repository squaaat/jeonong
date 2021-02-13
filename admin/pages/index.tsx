import { Fragment } from 'react'
import Link from 'next/link'

import styled, { createGlobalStyle } from 'styled-components'

import sg from 'utils/styleguide'

const IndexPage = () => (
  <Fragment>
    <GlobalLayout />
    <Layout>
      <SideNav>
        <SideNavHeader>
          조용진(Admin)
        </SideNavHeader>
        <SideNavContent>
          <ul>
            <li>
              <Link href="/products">상품관리</Link>
            </li>
          </ul>
        </SideNavContent>
      </SideNav>
      <Body>
        <Header>
          프로젝트 전농
        </Header>
        <Content>
          Content
        </Content>

      </Body>

    </Layout>
  </Fragment>
)

const Header = styled.div`
  width: 100%;
  padding: 0 0.5rem;

  height: ${sg.default.headerHeightSize};
  line-height: ${sg.default.headerHeightSize};

  background-color: ${sg.default.pointColor};
  color: ${sg.default.textColorR};
`

const Content = styled.div`
  width: 100%;
`
const SideNav = styled.nav`
  background-color: ${sg.default.pointWeekColor};

  width: ${sg.default.sideNavWidthSize};
`
const SideNavHeader = styled.div`
  padding: 0 0.5rem;
  height: ${sg.default.headerHeightSize};

  line-height: ${sg.default.headerHeightSize};
  font-size: ${sg.default.textSizeStrong};

  background-color: ${sg.default.pointStrongColor};
  color: ${sg.default.textColorR};
`

const SideNavContent = styled.div`
  line-height: ${sg.default.headerHeightSize};
  font-size: ${sg.default.textSize};
`

const Body = styled.div`
  width: 100%;
`

const Layout = styled.div`
  display: flex;
  flex-direction: row;

  min-width: 100vw;
  min-height: 100vh;
`

const GlobalLayout = createGlobalStyle`
  body {
    background-color: ${sg.default.backgroundColor};
    color: ${sg.default.textColor};
    margin: 0;
  }
`

export default IndexPage
