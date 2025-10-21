import resolveConfig from "tailwindcss/resolveConfig";
import tailwindConfig from "../../tailwind.config.js";
import { useAppContext } from "../context/AppContext.js";

const fullConfig = resolveConfig(tailwindConfig);
const colors = fullConfig.theme.colors;

export function useThemeColors() {
  const { isThemeDark } = useAppContext();

  return {
    primary: isThemeDark ? colors.dark["text-primary"] : colors.light.primary,

    header: isThemeDark
      ? colors.dark["text-primary"]
      : colors.light["text-inverted"],

    switch: {
      track: {
        true: colors.dark["text-secondary"],
        false: colors.dark["text-secondary"],
      },
      thumb: isThemeDark ? colors.dark["white"] : colors.light.primary,
    },
  };
}
