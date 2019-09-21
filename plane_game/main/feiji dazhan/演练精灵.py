import pygame
from plane_sprites import *
#游戏初始化
pygame.init()

#创建游戏的窗口 400 * 700
screen = pygame.display.set_mode((400,700))
#绘制背景图像
bg = pygame.image.load("./feiji/background.png")
#加载图像数据
#blit绘制图
screen.blit(bg,(0,0))
#update更新屏幕显示


#绘制英雄的飞机
hero = pygame.image.load("./feiji/hero1.png")
screen.blit(hero,(150,400))
pygame.display.update()

#创建时钟对象
clock = pygame.time.Clock()
hero_rect = pygame.Rect(150,450,102,126)

#创建敌机的精灵
enemy = GameSprite("./feiji/enemy0.png")
enemy1 = GameSprite("./feiji/enemy1.png",2)



#创建敌机的精灵组
enemy_group = pygame.sprite.Group(enemy,enemy1)





#游戏循环

while True:
    #可以指定飞机移动的频率
    clock.tick(60)
    #监听事件
    for event in pygame.event.get():
        #判断事件类型是否是退出事件
        if event.type == pygame.QUIT:
            print("游戏退出...")

            #调用quit方法
            #pygame.quit()

            #exit（）直接终止所有代码
            exit()

    #捕获事件
    event_list = pygame.event.get()
    if len(event_list) >0:
     print(event_list)
    #修改飞机的位置
    hero_rect.y -= 1
    #判断飞机的位置
    if hero_rect.bottom <= 0:
        hero_rect.y = 700
    #重新绘制背景图片
    screen.blit(bg,(0,0))
    #调用bilt方法绘制图片
    screen.blit(hero,hero_rect)
    #让精灵组调用两个方法
    #update   -让组中的所有精灵更新位置
    enemy_group.update()


    #draw   -  在screen上绘制所有的精灵
    enemy_group.draw(screen)



    #调用update方法更新显示
    pygame.display.update()





pygame.quit()