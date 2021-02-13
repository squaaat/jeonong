import { FC } from 'react'
import styled from 'styled-components'


type CardProps = {
  header?: FC
  children: FC
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
  color: #333;
  font-size: 1.25rem;
  height: ${defaultHeaderSize};
  line-height: ${defaultHeaderSize};
  border-bottom: 1px solid #acacac;
`

const CardLayout = styled.div`
  color: #333;
  border: 1px solid #acacac;
  border-radius: 6px;
  box-shadow: 3px 3px 3px #ccc;
`

export default Card
