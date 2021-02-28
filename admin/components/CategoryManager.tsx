import { FC, useState, useEffect, Dispatch } from 'react'
import { Card, Form, Select, Input, Radio } from 'antd';
import { Category, getCategories} from 'models/Category'

type CategoryManagerProps = {
}

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 18 },
};

type RequiredMark = boolean | 'optional';

const CategoryManager: FC<CategoryManagerProps> = () => {
  
  
  const [categories, setCategories] = useState(Array<Category>());
  const [isRefresh, setRefresh] = useState(false);

  useEffect(() => {
    getCategories().then((categories) => {
      console.log("is call once")
      setCategories(categories)
    })
  }, [])

  // const [form] = Form.useForm();
  // const [requiredMark, setRequiredMarkType] = useState<RequiredMark>('optional');

  // const onRequiredTypeChange = ({ requiredMarkValue }: { requiredMarkValue: RequiredMark }) => {
  //   setRequiredMarkType(requiredMarkValue);
  // };


  // console.log(form)
  // console.log(requiredMark)

  // useEffect(() => {
  //   if(isRefresh) {
  //     getCategories().then((categories) => {
  //       console.log("is call once")
  //       setCategories(categories)
  //       setRefresh(false)
  //     })
  //   }
  //   return function cleanup() {
  //     console.log("cleanup")
  //   };
  // }, [isRefresh]);


  const onCategorySubmit = (values: any) => {
    console.log('Success:', values);
    setRefresh(true)
  };

  const onCategorySubmitFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <Card title="제조사 관리">
      <Form
        {...layout}
        name="basic"
        initialValues={{ remember: true }}
        onFinish={onCategorySubmit}
        onValuesChange={onRequiredTypeChange}
        onFinishFailed={onCategorySubmitFailed}
      >
        <Form.Item label="name">
          <Input />
        </Form.Item>
        <Form.Item label="code">
          <Input />
        </Form.Item>
        <Form.Item label="depth" name="depth">
          <Radio.Group>
            <Radio.Button value="depth1">depth1</Radio.Button>
            <Radio.Button value="depth2">depth2</Radio.Button>
            <Radio.Button value="depth3">depth3</Radio.Button>
            <Radio.Button value="depth4">depth4</Radio.Button>
          </Radio.Group>
        </Form.Item>
        <Form.Item
          label="category"
          name="category"
          rules={[{ required: true, message: 'Please input your category!' }]}
        >
          <Select
            showSearch
            placeholder="type a category"
            optionFilterProp="children"
            onChange={(...rest) => console.log('category.onChange', ...rest)}
            onFocus={(...rest) => console.log('category.onFocus', ...rest)}
            onBlur={(...rest) => console.log('category.onBlur', ...rest)}
            onSearch={(...rest) => console.log('category.onSearch', ...rest)}
            filterOption={(input, option) =>
              option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
            }
          >
            <Select.Option key={0} value="jack">뚜레주루</Select.Option>
            <Select.Option key={1} value="lucy">뚜벅이</Select.Option>
            <Select.Option key={2} value="tom">또라이</Select.Option>
          </Select>
        </Form.Item>
      </Form>
    </Card>
  )
}
export default CategoryManager