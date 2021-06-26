import {createContext, ReactNode} from "react";

type Context = {
  Username: string | null;
  Room: string | null;
};

const context = createContext<Context>({
  Username: null,
  Room: null,
});

export default function GameContextProvider({children}: {children: ReactNode}) {
  
}