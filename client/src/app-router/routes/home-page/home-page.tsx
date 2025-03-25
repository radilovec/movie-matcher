import { RouteObject } from "react-router-dom";
import { ROUTES } from "../../../constants/routes";
import { Layout } from "../../../components/layout/Layout";
import { HomePage } from "../../../pages/home-page/HomePage";

export const homePage: RouteObject = {
  path: ROUTES.homePage,
  element: <Layout />,
  children: [
    {
      index: true,
      element: <HomePage />,
    },
  ],
};
