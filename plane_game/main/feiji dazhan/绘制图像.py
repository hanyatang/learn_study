import pygame

pygame.init()

#创建游戏的窗口 400 * 700
screen = pygame.display.set_mode((400,700))
#绘制背景图像
bg = pygame.image.load("./feiji/background.png")
#加载图像数据
#blit绘制图像
screen.blit(bg,(0,0))
#update更新屏幕显示
pygame.display.update()


while True:
    pass


pygame.quit()