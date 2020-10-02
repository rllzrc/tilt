import { MemoryRouter } from "react-router"
import React from "react"
import PathBuilder from "./PathBuilder"
import { createMemoryHistory } from "history"
import { mount } from "enzyme"
import { twoResourceView } from "./testdata"
import SidebarItem from "./SidebarItem"
import SidebarKeyboardShortcuts from "./SidebarKeyboardShortcuts"

const fakeHistory = createMemoryHistory()
const pathBuilder = new PathBuilder("localhost", "/")
const shortcuts = (items: SidebarItem[], selected: string) => {
  let root = mount(
    <MemoryRouter initialEntries={["/"]}>
      <SidebarKeyboardShortcuts
        items={items}
        selected={selected}
        history={fakeHistory}
        pathBuilder={pathBuilder}
      />
    </MemoryRouter>
  )
  return root
    .find(SidebarKeyboardShortcuts)
    .instance() as SidebarKeyboardShortcuts
}

it("navigates forwards", () => {
  let items = twoResourceView().resources.map(res => new SidebarItem(res))
  let sks = shortcuts(items, "")

  sks.onKeydown({ key: "j" } as KeyboardEvent)
  expect(fakeHistory.location.pathname).toEqual("/r/vigoda")
})

it("navigates forwards wrap", () => {
  let items = twoResourceView().resources.map(res => new SidebarItem(res))
  let sks = shortcuts(items, "snack")

  sks.onKeydown({ key: "j" } as KeyboardEvent)
  expect(fakeHistory.location.pathname).toEqual("/")
})

it("navigates backwards", () => {
  let items = twoResourceView().resources.map(res => new SidebarItem(res))
  let sks = shortcuts(items, "snack")

  sks.onKeydown({ key: "k" } as KeyboardEvent)
  expect(fakeHistory.location.pathname).toEqual("/r/vigoda")
})

it("navigates backwards wrap", () => {
  let items = twoResourceView().resources.map(res => new SidebarItem(res))
  let sks = shortcuts(items, "")

  sks.onKeydown({ key: "k" } as KeyboardEvent)
  expect(fakeHistory.location.pathname).toEqual("/r/snack")
})
