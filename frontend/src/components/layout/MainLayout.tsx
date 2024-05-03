import React from "react"

import SiteHeader from "./SiteHeader";

interface MainLayoutProps extends React.PropsWithChildren {}

export function MainLayout( {children}:MainLayoutProps) {

  return (
    <>
      <SiteHeader />
      {children}   
    </>
  );
}

