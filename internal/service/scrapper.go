package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/defaults"
	"github.com/go-rod/rod/lib/proto"

	"github.com/marianozunino/goashot/internal/dto"
	"github.com/marianozunino/goashot/internal/model"
	storage "github.com/marianozunino/goashot/internal/storage/json"
)

const cookieFile = "session.json"

type ScrapperService interface {
	Scrape() error
}

func registerScrapperService(repo storage.Repository) ScrapperService {
	return &scrapperService{
		repo: repo,
	}
}

type scrapperService struct {
	repo    storage.Repository
	browser *rod.Browser
	page    *rod.Page
}

// assert that Ine implements IneService
var _ ScrapperService = (*scrapperService)(nil)

func (s *scrapperService) Scrape() error {
	s.openBrowser()
	defer s.browser.MustClose()
	defer s.persistCookies()

	s.loadCookies()
	s.goToLogin()
	s.goToAshotSite()
	s.waitMenu()

	fmt.Println("Menu loaded, waiting for orders")
	time.Sleep(2 * time.Second)

	for _, order := range s.repo.GetOrders() {
		fmt.Println("Order found, picking menu item")
		// pick menu item
		s.pickMenuItem(order.OrderType)
		fmt.Println("Menu item picked, choosing extras")
		s.chooseExtras(order.Toppings)
		fmt.Println("Extras chosen, adding to cart")
		s.addNoteToOrder(order.User)
		s.confirmOrder()
		fmt.Println("Order added to cart")
	}
	// press any key to continue
	fmt.Println("Press any key to continue")
	var input string
	fmt.Scanln(&input)
	return nil
}

func (s *scrapperService) addNoteToOrder(user string) {
	// search textarea
	textarea := s.page.MustElement("textarea")
	textarea.MustEval(`()=>this.click()`)
	textarea.MustEval(`()=>this.scrollIntoView()`)
	textarea.Input(user)
}

func (s *scrapperService) confirmOrder() {
	ele := s.page.MustElementR("span", "Agregar y seguir comprando")
	ele.MustClick()
}

func (s *scrapperService) pickMenuItem(orderType dto.ShawarmaType) {
	product, _ := dto.Shawarmas.GetShawarma(orderType)
	productTag := fmt.Sprintf("[data-qa=product-info-%s]", product)
	element := s.page.MustElement(productTag)
	element.MustClick()
}

func (s *scrapperService) chooseExtras(toppings []model.Topping) {
	for _, topping := range toppings {
		toppingTag := fmt.Sprintf("div:has(> input#checkbox-%s)", topping.ID)
		element := s.page.MustElement(toppingTag)
		element.MustEval(`()=>this.click()`)
	}
}

func (s *scrapperService) openBrowser() {
	defaults.ResetWith("show")
	browser := rod.New().NoDefaultDevice().MustConnect()
	s.browser = browser
}

func (s *scrapperService) goToLogin() {
	s.page = s.browser.MustPage("https://www.rappi.com.uy/login?url_redirect=%2F&source=WEB_HEADER")
	s.page.MustWindowMaximize()
	// view port 1920x1080
	// s.page.MustSetViewport(1920, 1080, 1, false)
	s.page.MustWaitLoad()

	// if we are already logged in, we are going to get redirected to the home page
	// verify that we are in the login page
	if s.page.MustInfo().URL == "https://www.rappi.com.uy/login?url_redirect=%2F&source=WEB_HEADER" {
		s.doLogin()
	}

}

func (s *scrapperService) doLogin() {
	s.page.MustElement("#__next > div.styles__Container-sc-b1pk4f-0.dTBJvs > div.styles__MethodsContainer-sc-b1pk4f-5.irsykf > div.styles__LoginFormContainer-sc-foo08x-0.csiTcI > div:nth-child(2) > div:nth-child(1) > div > input[type=tel]").MustInput("########")
	s.page.MustElement("#__next > div.styles__Container-sc-b1pk4f-0.dTBJvs > div.styles__MethodsContainer-sc-b1pk4f-5.irsykf > div.styles__LoginFormContainer-sc-foo08x-0.csiTcI > div:nth-child(2) > div.styles__ButtonContainer-sc-foo08x-1.iuSeuU > button:nth-child(1)").MustClick()
	wait := s.page.EachEvent(func(e *proto.PageLoadEventFired) (stop bool) {
		return s.page.MustInfo().URL == "https://www.rappi.com.uy/"
	})
	wait()
}

func (s *scrapperService) goToAshotSite() {

	s.page.MustNavigate("https://www.rappi.com.uy/restaurantes/114497-ashot-shawarma")
	s.page.MustWaitLoad()
}

func (s *scrapperService) waitMenu() {
	productTag := fmt.Sprintf("[data-qa=product-info-%s]", dto.FalafelProductId)
	element := s.page.MustElement(productTag)
	element.MustWaitVisible()
}

func (s *scrapperService) loadCookies() {
	var cookies []*proto.NetworkCookie
	file, err := os.ReadFile(cookieFile)

	if err != nil {
		cookies = []*proto.NetworkCookie{}
	}

	err = json.Unmarshal(file, &cookies)
	if err != nil {
		cookies = []*proto.NetworkCookie{}
	}

	if len(cookies) > 0 {
		s.browser.MustSetCookies(cookies...)
	}

}

func (s *scrapperService) persistCookies() {
	cookies := s.browser.MustGetCookies()
	file, _ := json.MarshalIndent(cookies, "", " ")
	_ = os.WriteFile(cookieFile, file, 0644)
}
