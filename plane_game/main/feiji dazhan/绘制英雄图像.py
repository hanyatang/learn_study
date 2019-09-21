import pygame
#游戏初始化
pygame.init()

#创建游戏的窗口 400 * 700
screen = pygame.display.set_mode((400,700))
#绘制背景图像
bg = pygame.image.load("./feiji/background.png")
#加载图像数据
#blit绘制图像
screen.blit(bg,(0,0))
#update更新屏幕显示


#绘制英雄的飞机
hero = pygame.image.load("./feiji/hero1.png")
screen.blit(hero,(150,400))
pygame.display.update()

#创建时钟对象
clock = pygame.time.Clock()
#游戏循环
i = 1
while True:

    clock.tick(60)
    print(1)
    i += 1

    pass

pygame.quit()