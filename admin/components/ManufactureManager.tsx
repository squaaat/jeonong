import React, { FC } from 'react'
import {
  Card,
  Form,
  Input,
  notification,
  Button,
} from 'antd';
import { Manufacture, getManufactures, putManufacture } from 'store/models/Manufacture'


type ManufactureManagerProps = {
}

interface ManufactureFormData {
  code: string;
  name: string;
  companyRegistrationNumber: string;
}

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 18 },
};

const ManufactureManager: FC<ManufactureManagerProps> = () => {
  const [manufactures, setManufactures] = React.useState<Array<Manufacture>>([]);
  console.log('ManufactureManager.manufactures', manufactures)

  React.useEffect(() => {
    getManufactures().then((m) => setManufactures(m))
    return function cleanup() {
      console.log("cleanup")
    };
  }, []);


  const onManufactureSubmit = (m: ManufactureFormData) => {
    if (!m) return


    const manufacture: Manufacture = new Manufacture()
    putManufacture(manufacture).then((res) => {
      openNotification(
        'success',
        '제조업체 등록 성공',
        (<>
          <p>Name: {res.Name}</p>
          <p>Code: {res.Code}</p>
          <p>CompanyRegistrationNumber: {res.CompanyRegistrationNumber}</p>
        </>),
      )
      getManufactures().then((m) => setManufactures(m))
    }).catch((e) => {
      openNotification('error', "카테고리 관리등록 실패",  JSON.stringify(e))
    })
  };

  const onManufactureSubmitFailed = (errorInfo: any) => {
    openNotification('error', "카테고리 관리등록 실패", JSON.stringify(errorInfo))
  };

  return (
    <Card title="제조업체 정보 관리">
      <Form<ManufactureFormData>
        {...layout}
        name="basic"
        initialValues={{ remember: true }}
        onFinish={onManufactureSubmit}
        onFinishFailed={onManufactureSubmitFailed}
      >
        <Form.Item
          label="name"
          name="name"
          rules={[{ type: 'string', required: true, message: 'name 값을 입력해주세요.' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="code"
          name="code"
          rules={[{ type: 'string', required: true, message: 'code 값을 입력해주세요.' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="companyRegistrationNumber"
          name="companyRegistrationNumber"
          rules={[{ type: 'string', required: true, message: 'companyRegistrationNumber 값을 입력해주세요.' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: 6 }}>
          <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
            Submit
          </Button>
        </Form.Item>
      </Form>
    </Card>
  )
}

type NotiType = 'info' | 'success' | 'warn' | 'error' | 'open'
const openNotification = (notiType: NotiType, title: React.ReactNode, content: React.ReactNode) => {
  notification[notiType]({
    message: title,
    description: content,
    placement: "topRight",
  });
};

export default ManufactureManager