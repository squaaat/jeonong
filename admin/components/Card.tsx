import { FC, ReactNode } from 'react'
import styled from 'styled-components'

import sg from 'utils/styleguide'

type CardProps = {
  header?: any
  children: ReactNode
}

const Card: FC<CardProps> = ({ header, children }) => {

  return (
    <CardLayout>
      {
        header ? (<CardHeader>{header}</CardHeader>) : (null)
      }
      <CardBody
        theme={{
          isHeader:!!header,
        }}
      >
        {children}
      </CardBody>
    </CardLayout>
  )
}

const defaultHeaderSize = '3rem'

const CardBody = styled.div`
  width: calc(100% - 1rem);
  padding: 0.5rem;
 
  border-radius: 0px 0px 6px 6px;
`
CardBody.defaultProps = {
  theme: {
    isHeader: false,
  }
}

const CardHeader = styled.div`
  width: calc(100% - 1rem);
  padding: 0 0.5rem;

  display: flex;
  flex-direction: row;

  color: ${sg.default.textStrongColor};
  font-size: ${sg.default.textSizeTitle};
  height: ${defaultHeaderSize};
  line-height: ${defaultHeaderSize};
  border-bottom: ${sg.default.componentBorderSolid} solid ${sg.default.componentBorderColor};
`

const CardLayout = styled.div`
  color: ${sg.default.textColor};
  font-size: ${sg.default.textSize};

  border: ${sg.default.componentBorderSolid} solid ${sg.default.componentBorderColor};
  background-color: ${sg.default.componentBackgroundColor};
  border-radius: 6px;
  box-shadow: ${sg.default.componentBoxShadowSize} ${sg.default.componentBoxShadowSize} ${sg.default.componentBoxShadowSize} ${sg.default.componentBoxShadowColor};
`

export default Card
