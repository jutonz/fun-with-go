package newsfeed

import "testing"

func TestAdd(t *testing.T) {
  feed := New()
  feed.Add(Item{
    Title: "title",
    Post: "post",
  })

  if len(feed.Items) != 1 {
    t.Error("item was not added")
  }
}

func TestGetAll(t *testing.T) {
  feed := New()
  feed.Add(Item{
    Title: "title",
    Post: "post",
  })

  result := feed.GetAll()

  if len(result) != 1 {
    t.Error("not all items were returned")
  }
}
