import oc from 'open-color'

type Guides = {
  default: Guide;
}

type Guide = {
  textColor: string;
  textStrongColor: string;
  textWeekColor: string;

  textColorR: string;
  textStrongColorR: string;
  textWeekColorR: string;

  textSize: string;
  textSizeTitle: string;
  textSizeWeek: string;
  textSizeStrong: string;

  textWeight: number;
  textWeightStrong: number;
  textWeightWeek: number;

  componentBorderSolid: string;
  componentBorderColor: string;
  componentBorderRadius: string;

  primaryColor: string;
  primaryStrongColor: string;
  primaryWeekColor: string;

  pointColor: string;
  pointStrongColor: string;
  pointWeekColor: string;

  backgroundColor: string;

  headerHeightSize: string;
  sideNavWidthSize: string;
}

const defaultColor: Guide = {
  textColor: oc.gray[8],
  textStrongColor: oc.black,
  textWeekColor: oc.gray[6],

  textColorR: oc.gray[1],
  textStrongColorR: oc.white,
  textWeekColorR: oc.gray[3],

  textSize: '1rem',
  textSizeTitle: '1.25rem',
  textSizeWeek: '0.75rem',
  textSizeStrong: '1.125rem',

  textWeight: 400,
  textWeightStrong: 700,
  textWeightWeek: 300,

  componentBorderSolid: '1px',
  componentBorderColor: oc.gray[4],
  componentBorderRadius: '5px',
  
  primaryColor: oc.gray[6],
  primaryStrongColor: oc.gray[8],
  primaryWeekColor: oc.gray[4],

  pointColor: oc.pink[6],
  pointStrongColor: oc.pink[9],
  pointWeekColor: oc.pink[2],

  backgroundColor: oc.gray[0],

  headerHeightSize: '3rem',

  sideNavWidthSize: '20rem',
}
const styleGuides: Guides = {
  default: defaultColor,
}
export default styleGuides