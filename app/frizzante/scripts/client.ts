import { hydrate } from "svelte"
import ClientView from "../components/ClientView.svelte"
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
hydrate(ClientView, { target: target(), props: props() })
