import {
  create,
  NAvatar,
  NButton,
  NLayout,
  NInput,
  NTag,
} from "naive-ui";

const naive = create({
  components: [
    NAvatar,
    NButton,
    NLayout,
    NInput,
    NTag,
  ],
})

export function setUpNaiveUI(app) {
  app.use(naive)
}